import Link from "@/components/Link";
import MainLayout from "@/layouts/main";
import { getAuthToken } from "@/util/auth";
import { Box, Button, CopyButton, List, Title } from "@mantine/core";
import { useEffect, useState } from "react";
import { useJwt } from "react-jwt";

function ExportPage() {
  const { decodedToken } = useJwt(getAuthToken() || "");
  const [id, setId] = useState<string | null>(null);

  useEffect(() => {
    if (decodedToken) {
      const jwt = decodedToken as Record<string, string>;
      setId(jwt["id"]);
    }
  }, [decodedToken]);

  return (
    <MainLayout>
      <Box className="px-8 pt-8 sm:px-16">
        <Title order={1}>Installing your calendar</Title>
        <CopyButton value={`https://uogcal.ronkzd.xyz/calendar/${id}`}>
          {({ copied, copy }) => (
            <Button onClick={copy}>{copied ? "Copied" : "Copy Link"}</Button>
          )}
        </CopyButton>
        <List listStyleType="disc">
          <List.Item>
            Using Google Calendar
            <List listStyleType="disc">
              <List.Item>
                Follow the steps under the 'Use a link to add a public calendar'
                section of{" "}
                <Link
                  to="https://support.google.com/calendar/answer/37100"
                  target="_blank"
                  className="underline"
                  rel="noopener noreferrer"
                >
                  {"this page."}
                </Link>
              </List.Item>
            </List>
          </List.Item>
          <List.Item>
            <Link
              to="http://www.medportal.ca/public/help/google/calendar/import-a-calendar-by-url-on-iphone-ipad-ipod"
              target="_blank"
              className="underline"
              rel="noopener noreferrer"
            >
              Using the Apple calendar app
            </Link>
          </List.Item>
          <li>
            Use a calendar not listed here? Google "how to add a calendar by
            url" followed by your calendar app's name
          </li>
        </List>
      </Box>
    </MainLayout>
  );
}

export default ExportPage;
