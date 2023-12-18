package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var instance *pgxpool.Pool

func Setup() {
	// Warm up db
	getInstance()
}

// This is a private function to keep all the SQL in one place
func getInstance() *pgxpool.Pool {
	if instance == nil {
		instance = connect()
	}

	return instance
}

func connect() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%v:%v@%v/%v",
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_DATABASE"),
	))
	if err != nil {
		panic("Couldn't create db config with given env variables")
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic("Couldn't connect to db (check if env vars are correct)")
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		panic("Couldn't ping the db")
	}

	return dbpool
}

// Add a new user to db
func AddUser(displayName string, password string) (uuid.UUID, error) {
	pool := getInstance()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	userUUID := uuid.Must(uuid.NewUUID())

	_, err = pool.Exec(context.Background(), `INSERT INTO uogcal_user VALUES($1, $2, $3);`, userUUID, displayName, string(hashedPassword))
	if err != nil {
		return uuid.Nil, err
	}
	return userUUID, nil
}

func AddOAuthUser(oauthId string) (string, uuid.UUID, error) {
	pool := getInstance()

	row := pool.QueryRow(context.Background(), "SELECT 1 FROM oauth_user ou WHERE ou.user_uid=$1;", oauthId)
	var res string
	if row.Scan(&res) == nil {
		return "", uuid.Nil, fmt.Errorf("user already exists. can't register")
	}

	userUUID := uuid.Must(uuid.NewUUID())

	username := userUUID.String()
	_, err := pool.Exec(context.Background(), `INSERT INTO uogcal_user VALUES($1, $2, $3);`, userUUID, username, "")
	if err != nil {
		return "", uuid.Nil, err
	}

	_, err = pool.Exec(context.Background(), `INSERT INTO oauth_user VALUES($1, $2);`, oauthId, userUUID)
	if err != nil {
		return "", uuid.Nil, err
	}

	return username, userUUID, nil
}

// Add a section+meetings
func AddSection(section CourseSection, meetings []SectionMeeting) error {
	pool := getInstance()

	// Insert course into table of sections
	_, err := pool.Exec(
		context.Background(),
		`INSERT INTO course_section VALUES($1, $2)`,
		section.Code,
		section.Name,
	)
	if err != nil {
		return err
	}

	// Insert sections
	for _, meeting := range meetings {
		_, err = pool.Exec(
			context.Background(),
			`INSERT INTO section_meeting VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			meeting.Code,
			meeting.Type,
			fmt.Sprintf(
				"%v %v.%v",
				meeting.Created.Format(time.DateOnly),
				meeting.Created.Format(time.TimeOnly),
				meeting.Created.Nanosecond()%1e6,
			),
			meeting.StartDate.Format(time.DateOnly),
			meeting.EndDate.Format(time.DateOnly),
			meeting.StartTime.Format(time.TimeOnly),
			meeting.EndTime.Format(time.TimeOnly),
			meeting.MeetingDays,
			meeting.Location,
			fmt.Sprintf(
				"%v %v.%v",
				meeting.LastModified.Format(time.DateOnly),
				meeting.LastModified.Format(time.TimeOnly),
				meeting.LastModified.Nanosecond()%1e6,
			),
			meeting.UpdateCount,
		)
		if err != nil {
			// TODO rollback the new course_section entry on error
			return err
		}
	}

	return nil
}

func AddNewSections(section CourseSection, meetings []SectionMeeting) error {
	pool := getInstance()

	// Insert sections
	for _, meeting := range meetings {
		_, err := pool.Exec(
			context.Background(),
			`INSERT INTO section_meeting VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT DO NOTHING`,
			meeting.Code,
			meeting.Type,
			fmt.Sprintf(
				"%v %v.%v",
				meeting.Created.Format(time.DateOnly),
				meeting.Created.Format(time.TimeOnly),
				meeting.Created.Nanosecond()%1e6,
			),
			meeting.StartDate.Format(time.DateOnly),
			meeting.EndDate.Format(time.DateOnly),
			meeting.StartTime.Format(time.TimeOnly),
			meeting.EndTime.Format(time.TimeOnly),
			meeting.MeetingDays,
			meeting.Location,
			fmt.Sprintf(
				"%v %v.%v",
				meeting.LastModified.Format(time.DateOnly),
				meeting.LastModified.Format(time.TimeOnly),
				meeting.LastModified.Nanosecond()%1e6,
			),
			meeting.UpdateCount,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

// Add a user to course
func AddUserToSection(userId string, code string) error {
	pool := getInstance()

	_, err := pool.Exec(
		context.Background(),
		`INSERT INTO uogcal_user_course_section VALUES($1, $2)`,
		userId,
		code,
	)

	return err
}

func GetUserPassword(username string) (string, string, error) {
	pool := getInstance()
	row := pool.QueryRow(
		context.Background(),
		"SELECT uid, password FROM uogcal_user u WHERE u.display_name=$1;",
		username,
	)
	var uid string
	var password []byte
	err := row.Scan(&uid, &password)

	return uid, string(password), err
}

func GetOAuthUser(oauthId string) (string, string, error) {
	pool := getInstance()
	row := pool.QueryRow(
		context.Background(),
		`SELECT display_name, user_uid FROM oauth_user ou JOIN uogcal_user uu ON ou.user_uid=uu.uid WHERE oauth_id=$1;`,
		oauthId,
	)
	var username string
	var uid string
	err := row.Scan(&username, &uid)

	return username, uid, err
}

func CheckCourseExists(code string) bool {
	pool := getInstance()

	res, err := pool.Exec(
		context.Background(),
		`SELECT 1 FROM course_section cs WHERE cs.code=$1`,
		code,
	)
	if err != nil {
		return false
	}
	return res.RowsAffected() == 1
}

// Get sections by uuid
func GetSectionsByUUID(uuid string) ([]CourseSectionJoin, error) {
	pool := getInstance()

	rows, err := pool.Query(
		context.Background(),
		`SELECT
            cs.code,
            cs.name,
            sm.type,
			sm.created,
            sm.start_date,
            sm.end_date,
            sm.start_time,
            sm.end_time,
            sm.meeting_days,
            sm.location,
			sm.last_modified,
			sm.update_count
        FROM
            section_meeting sm
        INNER JOIN
            course_section cs ON sm.code=cs.code
        INNER JOIN
            uogcal_user_course_section u_cs ON u_cs.course_section_code=cs.code
        INNER JOIN
            uogcal_user u ON u.uid=u_cs.uogcal_user_uid
        WHERE u.uid=$1;`,
		uuid,
	)
	if err != nil {
		return nil, err
	}

	var res []CourseSectionJoin
	for rows.Next() {
		var name string
		var meeting SectionMeeting

		err := rows.Scan(
			&meeting.Code,
			&name,
			&meeting.Type,
			&meeting.Created,
			&meeting.StartDate,
			&meeting.EndDate,
			&meeting.StartTime,
			&meeting.EndTime,
			&meeting.MeetingDays,
			&meeting.Location,
			&meeting.LastModified,
			&meeting.UpdateCount,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, CourseSectionJoin{
			Name:    name,
			Meeting: meeting,
		})
	}
	return res, nil
}

func GetSectionsBeforeDate(date time.Time, term string) (map[string]bool, error) {
	pool := getInstance()

	rows, err := pool.Query(
		context.Background(),
		"SELECT code FROM section_meeting sm WHERE sm.created <= $1;",
		date,
	)
	if err != nil {
		return nil, err
	}

	alreadyChecked := map[string]bool{}

	for rows.Next() {
		var code string

		err := rows.Scan(&code)
		if err != nil {
			return nil, err
		}
		if alreadyChecked[code] {
			continue
		}

		alreadyChecked[code] = true
	}

	return alreadyChecked, nil
}
