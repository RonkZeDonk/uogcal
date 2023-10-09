import { useNavigate } from "react-router-dom";
import {
  Button,
  Center,
  Grid,
  Group,
  Space,
  Stack,
  Text,
  Title,
} from "@mantine/core";
import LandingHero from "@/components/LandingHero";
import MainLayout from "@/layouts/main";
import { isLoggedIn } from "@/util/auth";

function Index() {
  const navigate = useNavigate();

  return (
    <>
      <MainLayout>
        <Center mih="calc(100vh - 4.5rem)" className="px-8 sm:px-16">
          <Grid>
            <Grid.Col span={{ base: 12, sm: 6 }} className="max-sm:text-center">
              <Stack justify="center" h="100%" mr="lg">
                <Title order={1} fw="600">
                  Save time without sacrificing your experience
                </Title>
                <Text>
                  Import all your courses in a couple clicks. Never forget about
                  an assignment with our classroom-wide assignment date sharing.
                </Text>
                <Group className="max-sm:self-center">
                  <Button
                    size="md"
                    onClick={() => !isLoggedIn() && navigate("/signup")}
                    variant="filled"
                  >
                    Get Started
                  </Button>
                  <Button
                    size="md"
                    variant="subtle"
                    onClick={() => !isLoggedIn() && navigate("/login")}
                  >
                    {/* TODO create a learn more button or something of the sorts */}
                    Returning user
                  </Button>
                </Group>
                <Space h="xl" />
              </Stack>
            </Grid.Col>
            <Grid.Col span={{ base: 12, sm: 6 }}>
              <Center h="100%">
                <LandingHero />
              </Center>
            </Grid.Col>
          </Grid>
        </Center>
      </MainLayout>
    </>
  );
}

export default Index;
