import ImportSection from "@/components/ImportSection";
import Link from "@/components/Link";
import MainLayout from "@/layouts/main";
import {
  Button,
  Code,
  Group,
  List,
  Space,
  Text,
  Title,
  rem,
} from "@mantine/core";
import { Dropzone, FileWithPath } from "@mantine/dropzone";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { MdCloudDone, MdCloudUpload } from "react-icons/md";

import "@mantine/dropzone/styles.css";
import { IconContext } from "react-icons";

const BOOKMARKLET_LOCATION = "/files/getCourses.js";

function ImportPage() {
  const navigate = useNavigate();
  const [bookmarklet, setBookmarklet] = useState<string>("");
  const loading = bookmarklet === "";

  useEffect(() => {
    fetch(BOOKMARKLET_LOCATION).then((res) => {
      res.text().then((js) => {
        // Double scoped to keep the window object clean
        setBookmarklet(`javascript:(()=>{ { ${encodeURIComponent(js)} } })()`);
      });
    });
  }, []);

  const onDrop = (files: FileWithPath[]) => {
    files[0].text().then((text) => {
      const jsonData = JSON.parse(text);

      const fd = new FormData();
      for (const key in jsonData) {
        fd.append(key, jsonData[key]);
      }
      fetch("/upload/courses", { method: "post", body: fd });
    });

    const data = files[0];
    const fd = new FormData();
    fd.append("courses", data);

    fetch("upload/courses", { method: "post", body: fd }).then(() => navigate("/account"));
  };

  return (
    <MainLayout>
      <div className="px-8 py-12 sm:px-16">
        <Title order={1}>Import your courses</Title>

        <Text>
          Entering your courses into a calendar manually is a thing of the past.
          Now you can use a{" "}
          <Link
            to="https://en.wikipedia.org/wiki/Bookmarklet"
            target="_blank"
            className="underline"
          >
            {"bookmarklet"}
          </Link>{" "}
          to quickly export your courses into a json file. We'll take that json
          file and import it on to your account and automatically manage your
          course schedule.
        </Text>

        <ImportSection title="Note:">
          <Text>
            The bookmarklet provided doesn't access your account credentials,
            storage, or the internet. It simply scrapes the page for your course
            info.
          </Text>
          <Text>
            If your schedule changes at any time (for example you drop/add a
            course) you need to update your courses manually by rerunning the
            bookmarklet and reuploading the json data.
          </Text>
        </ImportSection>

        <Text className="mb-2">
          Drag this bookmarklet to your bookmarks bar.
        </Text>
        <Space h="1rem" />

        <Button
          loading={loading}
          loaderProps={{ type: "dots" }}
          component={loading ? "div" : "a"}
          href={bookmarklet}
          variant="filled"
          className="text-inherit hover:text-inherit"
        >
          Course Exporter
        </Button>
        <Button variant="subtle" onClick={() => navigate(BOOKMARKLET_LOCATION)}>
          View bookmarklet source
        </Button>

        <Text>
          Once you have the bookmarklet installed you can use it to generate
          your json data in an easy 3 step process.
        </Text>
        <Space h="0.5rem" />
        <List listStyleType="numeric" withPadding>
          <List.Item>
            Go to the{" "}
            <Link
              to="https://colleague-ss.uoguelph.ca/Student/Planning/DegreePlans"
              target="_blank"
              className="underline"
            >
              plan and schedule page
            </Link>{" "}
            on WebAdvisor and navigate to your target semester.
          </List.Item>
          <List.Item>
            Launch the bookmarklet and save the file it generates.
          </List.Item>
          <List.Item>Import the file below</List.Item>
        </List>

        <Space my="md" />

        <Dropzone
          onDrop={onDrop}
          onReject={() => {}}
          maxSize={1024 * 5} // 5 KiB file limit
          maxFiles={1}
          accept={["application/json"]}
        >
          <Group
            justify="center"
            gap="xl"
            mih={220}
            style={{ pointerEvents: "none" }}
          >
            <Dropzone.Accept>
              <IconContext.Provider
                value={{
                  size: rem(52),
                  color: "var(--mantine-color-dimmed)",
                }}
              >
                <MdCloudDone />
              </IconContext.Provider>
            </Dropzone.Accept>
            <Dropzone.Reject>
              <IconContext.Provider
                value={{
                  size: rem(52),
                  color: "var(--mantine-color-dimmed)",
                }}
              >
                <MdCloudDone />
              </IconContext.Provider>
            </Dropzone.Reject>
            <Dropzone.Idle>
              <IconContext.Provider
                value={{
                  size: rem(52),
                  color: "var(--mantine-color-dimmed)",
                }}
              >
                <MdCloudUpload />
              </IconContext.Provider>
            </Dropzone.Idle>

            <div>
              <Text size="xl" inline>
                Drag file here or click to select file
              </Text>
              <Text size="sm" c="dimmed" inline mt={7}>
                Upload the <Code>courses.json</Code> here. Should not exceed 5
                KiB
              </Text>
            </div>
          </Group>
        </Dropzone>
      </div>
    </MainLayout>
  );
}

export default ImportPage;
