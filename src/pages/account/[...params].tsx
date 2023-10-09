import Link from "@/components/Link";
import MainLayout from "@/layouts/main";
import { Box, Button, ButtonGroup, Divider, Flex, List, ListItem, Space, Stack, Tabs, Text, Title } from "@mantine/core";
import { DatePicker } from "@mantine/dates";
import { useEffect, useState } from "react";
import { IconContext } from "react-icons";
import { TbCalendarShare } from "react-icons/tb";
import { useNavigate, useParams } from "react-router-dom";

interface SectionMeeting {
  code: string
  type: string
  created: string
  startDate: string
  endDate: string
  startTime: string
  endTime: string
  meetingDays: number[]
  location: string
  lastModified: string
  updateCount: number
}
interface CourseData {
  name: string
  sectionMeeting: SectionMeeting
}

const idxToDate = (idx: number) => {
  return [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ][idx];
};

const toClockTime = (d: Date) => {
  const [h, m] = [d.getUTCHours(), d.getUTCMinutes()];
  return `${h % 12}:${m}${h/12 < 1 ? "am" : "pm"}`;
};

function AccountPage() {
  const params = useParams();
  const pageCode = params["*"] ? params["*"] : null;

  const navigate = useNavigate();

  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date(Date.now()));
  const [courses, setCourses] = useState<CourseData[] | null>(null);
  const [unique, setUnique] = useState<JSX.Element[] | null>(null);
  const [idx, setIdx] = useState(0);
  useEffect(() => {
    fetch("/api/courses").then((data) =>
      data.json().then((courseData: CourseData[]) => {
        if (!courseData) {
          setCourses([]);
          setUnique([]);

          return;
        }

        const uq: string[] = [];
        let i = 0;
        for (const { name, sectionMeeting } of courseData) {
          const key = name + ":" + sectionMeeting.code;
          if (sectionMeeting.type === "LEC" && sectionMeeting.code === pageCode)
            setIdx(i);
          if (uq.indexOf(key) === -1) uq.push(key);
          i++;
        }

        setCourses(courseData);
        setUnique(uq.map((title) => {
          const arr = title.split(":");
          const code = arr.pop();
          const name = arr.join(":");
          return (
            <Link to={`/account/${code}`} className="underline">
              {name}
            </Link>
          );
        }));
      })
    );
  }, [pageCode]);

  return (
    <MainLayout>
      <Flex p="md" gap="md" className="h-[calc(100vh-4.5rem)] w-inherit">
        <Flex className="w-72">
          <Stack
            gap="md"
            className="[&>*]:border-2 [&>*]:rounded-md [&>*]:border-[var(--mantine-color-default-border)]"
          >
            <DatePicker
              defaultValue={selectedDate}
              onChange={(e) => {
                setSelectedDate(e);
              }}
              p="md"
              firstDayOfWeek={0}
            />
            <Box p="lg" className="flex-1 overflow-auto">
              {/* TODO Use events as the default */}
              {/* <Tabs defaultValue="events"> */}
              <Tabs defaultValue="courses">
                <Tabs.List>
                  <Tabs.Tab value="courses">Your Courses</Tabs.Tab>
                  {/* <Tabs.Tab value="events">Upcoming events</Tabs.Tab> */}
                </Tabs.List>
                <Space h="8px" />

                <Tabs.Panel value="events">
                  <Text>Next 10 events</Text>
                  <List withPadding listStyleType="circle">
                    <List.Item>Coming Soon</List.Item>
                  </List>
                </Tabs.Panel>

                <Tabs.Panel value="courses">
                  <List listStyleType="disc">
                    {unique
                      ? unique.map((name) => <ListItem>{name}</ListItem>)
                      : "loading"}
                  </List>
                </Tabs.Panel>
              </Tabs>
            </Box>
          </Stack>
        </Flex>
        <Box
          p="xl"
          className="flex-1 border-2 rounded-md border-[var(--mantine-color-default-border)]"
        >
          <Flex h="max-content">
            <Title order={1}>{courses ? courses[idx].name : "Loading"}</Title>
            <Title
              order={2}
              size="1.2rem"
              c="var(--mantine-color-dimmed)"
              my="auto"
              ml="xs"
              h="100%"
            >
              {courses ? " — " + courses[idx].sectionMeeting.location : "loading"}
            </Title>
            <Flex justify="end" className="flex-1">
              <ButtonGroup>
                <Button onClick={() => navigate("/export")}>
                  <IconContext.Provider value={{ size: "18px" }}>
                    <TbCalendarShare />
                  </IconContext.Provider>
                  <Text ml="xs">Export Calendar</Text>
                </Button>
              </ButtonGroup>
            </Flex>
          </Flex>
          <Text size="sm">
            {courses
              ? courses[idx].sectionMeeting.meetingDays
                  .map((day) => idxToDate(day))
                  .join(", ") +
                " @ " +
                toClockTime(new Date(courses[0].sectionMeeting.startTime))
              : "Loading"}
          </Text>
          <Divider my="md" />
          <List listStyleType="disc" withPadding>
            <Text>Upcoming events</Text>
            <ListItem>This feature is coming soon</ListItem>
          </List>
          <Divider my="md" />
          {/* <Text>Two other users have this course</Text> */}
        </Box>
      </Flex>
    </MainLayout>
  );
}

export default AccountPage;
