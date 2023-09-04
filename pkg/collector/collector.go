package collector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/RonkZeDonk/uogcal/pkg/database"
)

// Sources of data
//  * Course schedule: https://colleague-ss.uoguelph.ca/Student/Student/Courses
//  * Building codes: https://www.uoguelph.ca/registrar/scheduling/buildingcodes-col

type searchResults struct {
	Courses []struct {
		MatchingSectionIds []string
		Id                 string
		SubjectCode        string
		Number             string
	}
}

type sectionResults struct {
	TermsAndSections []struct {
		Sections []struct {
			Section struct {
				SectionNameDisplay  string
				SectionTitleDisplay string
				TermId              string

				FormattedMeetingTimes []struct {
					InstructionalMethodDisplay string
					BuildingDisplay            string
					RoomDisplay                string
					DatesDisplay               string
					StartTime                  string
					EndTime                    string
					Days                       []int8
				}
			}
		}
	}
}

type Verification struct {
	Header string
	Cookie string
}

func setHeaders(r *http.Request, verifyHeader string, verifyCookie string) {
	r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/114.0")
	r.Header.Add("Accept", "application/json; q=0.01")
	r.Header.Add("Content-Type", "application/json, charset=utf-8")
	r.Header.Add("X-Requested-With", "XMLHttpRequest")
	r.Header.Add("Origin", "https://colleague-ss.uoguelph.ca")
	r.Header.Add("DNT", "1")
	r.Header.Add("Connection", "keep-alive")

	r.Header.Add("__RequestVerificationToken", verifyHeader)
	r.Header.Add("Cookie", verifyCookie)
}

func GetVerificationTokenAndCookie() (Verification, error) {
	verify := Verification{}

	res, err := http.Get("https://colleague-ss.uoguelph.ca/Student/Courses")
	if err != nil {
		return Verification{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return Verification{}, errors.New("response wasn't successful (>299)")
	}
	if err != nil {
		return Verification{}, err
	}

	r, err := regexp.Compile(`<input.*value="([-_a-zA-Z0-9]*)".*\/>`)
	if err != nil {
		return Verification{}, err
	}
	verify.Header = r.FindStringSubmatch(string(body))[1]

	for _, cookie := range res.Cookies() {
		verify.Cookie = fmt.Sprintf("%v=%v", cookie.Name, cookie.Value)
	}

	return verify, nil
}

func GetSectionData(term string, sectionCode string, verifyHeader string, verifyCookie string) (database.CourseSection, []database.SectionMeeting, error) {
	courseCode := strings.Join(strings.Split(sectionCode, "*")[:2], "*")

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		"https://colleague-ss.uoguelph.ca/Student/Courses/SearchAsync",
		strings.NewReader(
			fmt.Sprintf(
				`{"searchParameters":"{\"keyword\":\"%v\",\"terms\":[],\"requirement\":null,\"subrequirement\":null,\"courseIds\":null,\"sectionIds\":null,\"requirementText\":null,\"subrequirementText\":\"\",\"group\":null,\"startTime\":null,\"endTime\":null,\"openSections\":null,\"subjects\":[],\"academicLevels\":[],\"courseLevels\":[],\"synonyms\":[],\"courseTypes\":[],\"topicCodes\":[],\"days\":[],\"locations\":[],\"faculty\":[],\"onlineCategories\":null,\"keywordComponents\":[],\"startDate\":null,\"endDate\":null,\"startsAtTime\":null,\"endsByTime\":null,\"pageNumber\":1,\"sortOn\":\"None\",\"sortDirection\":\"Ascending\",\"subjectsBadge\":[],\"locationsBadge\":[],\"termFiltersBadge\":[],\"daysBadge\":[],\"facultyBadge\":[],\"academicLevelsBadge\":[],\"courseLevelsBadge\":[],\"courseTypesBadge\":[],\"topicCodesBadge\":[],\"onlineCategoriesBadge\":[],\"openSectionsBadge\":\"\",\"openAndWaitlistedSectionsBadge\":\"\",\"subRequirementText\":null,\"quantityPerPage\":30,\"openAndWaitlistedSections\":null,\"searchResultsView\":\"CatalogListing\"}"}`,
				sectionCode,
			),
		),
	)
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}

	setHeaders(req, verifyHeader, verifyCookie)
	res, err := client.Do(req)
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}
	if res.StatusCode != 200 {
		return database.CourseSection{}, []database.SectionMeeting{}, errors.New("status was not 200")
	}

	bodyRaw, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}
	var results searchResults
	err = json.Unmarshal(bodyRaw, &results)
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}

	var courseId string
	var sectionIds []string
	for _, result := range results.Courses {
		resultCode := result.SubjectCode + "*" + result.Number
		if !strings.EqualFold(courseCode, resultCode) {
			continue
		}

		courseId = result.Id
		sectionIds = append(sectionIds, result.MatchingSectionIds...)

		break
	}

	req, err = http.NewRequest(
		"POST",
		"https://colleague-ss.uoguelph.ca/Student/Courses/SectionsAsync",
		strings.NewReader(
			fmt.Sprintf(
				`{"courseId":"%v","sectionIds":[%v]}`,
				courseId,
				strings.Join(sectionIds, ","),
			),
		),
	)
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}

	setHeaders(req, verifyHeader, verifyCookie)
	res, err = client.Do(req)
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}
	if res.StatusCode != 200 {
		return database.CourseSection{}, []database.SectionMeeting{}, errors.New("status was not 200")
	}

	bodyRaw, err = io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}

	var sectionResults sectionResults
	err = json.Unmarshal(bodyRaw, &sectionResults)
	if err != nil {
		return database.CourseSection{}, []database.SectionMeeting{}, err
	}

	var section database.CourseSection
	var meetings []database.SectionMeeting
	for _, ts := range sectionResults.TermsAndSections {
		for _, s := range ts.Sections {
			if s.Section.TermId != term {
				continue
			}
			section.Code = s.Section.SectionNameDisplay
			section.Name = s.Section.SectionTitleDisplay

			for _, meeting := range s.Section.FormattedMeetingTimes {
				if meeting.DatesDisplay == "" {
					continue
				}
				dates := strings.Split(meeting.DatesDisplay, " - ")
				startDate, err := time.Parse("1/2/2006", dates[0])
				if err != nil {
					return database.CourseSection{}, []database.SectionMeeting{}, err
				}
				endDate, err := time.Parse("1/2/2006", dates[1])
				if err != nil {
					return database.CourseSection{}, []database.SectionMeeting{}, err
				}

				if meeting.StartTime == "" {
					continue
				}
				startTime, err := time.Parse(time.TimeOnly, meeting.StartTime)
				if err != nil {
					return database.CourseSection{}, []database.SectionMeeting{}, err
				}
				endTime, err := time.Parse(time.TimeOnly, meeting.EndTime)
				if err != nil {
					return database.CourseSection{}, []database.SectionMeeting{}, err
				}

				now := time.Now()
				meetings = append(meetings, database.SectionMeeting{
					Code:         sectionCode,
					Type:         meeting.InstructionalMethodDisplay,
					Created:      now,
					StartDate:    startDate,
					EndDate:      endDate,
					StartTime:    startTime,
					EndTime:      endTime,
					MeetingDays:  meeting.Days,
					Location:     fmt.Sprintf("%v %v", meeting.BuildingDisplay, meeting.RoomDisplay),
					LastModified: now,
					UpdateCount:  0,
				})
			}
		}
	}

	return section, meetings, nil
}

func AddNewSections(date time.Time, term string) error {
	matchedSections, err := database.GetSectionsBeforeDate(date, term)
	if err != nil {
		return err
	}

	verify, err := GetVerificationTokenAndCookie()
	if err != nil {
		return err
	}

	for code := range matchedSections {
		course, meetings, err := GetSectionData(strings.ToUpper(term), code, verify.Header, verify.Cookie)
		if err != nil {
			return err
		}
		err = database.AddNewSections(course, meetings)
		if err != nil {
			return err
		}
	}

	fmt.Printf("checked %v courses\n", len(matchedSections))

	return nil
}
