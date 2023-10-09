import Link from "@/components/Link";
import { isLoggedIn, logout } from "@/util/auth";
import { Center, Flex, Group, Text, UnstyledButton } from "@mantine/core";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

interface MainLayoutProps {
  children?: string | string[] | JSX.Element | JSX.Element[];
}


function MainLayout({ children }: MainLayoutProps) {
  const navigate = useNavigate();

  const [loggedIn, setLoggedIn] = useState<boolean>(isLoggedIn());

  const logoutOnClick = () => {
    logout();
    setLoggedIn(false);
    navigate("/");
  };

  return (
    <div className="h-full">
      <Flex
        p="static"
        w="100%"
        h="4.5rem"
        justify="space-between"
        px="4rem"
        className="bg-[var(--mantine-color-gray-4)] dark:bg-[var(--mantine-color-dark-8)]"
      >
        <Center>
          <Link to="/">
            <Text size="1.4rem">UoG Calendar</Text>
          </Link>
        </Center>

        <Group gap="xl">
          <Link className="max-sm:hidden" to="/">
            Home
          </Link>
          {!loggedIn && (
            <>
              <Link className="max-sm:hidden" to="/login">
                Log In
              </Link>
              <Link className="max-sm:hidden" to="/signup">
                Sign Up
              </Link>
            </>
          )}
          {!!loggedIn && (
            <>
              <Link className="max-sm:hidden" to="/import">
                Import
              </Link>
              <Link className="max-sm:hidden" to="/account">
                Account
              </Link>
              <UnstyledButton className="max-sm:hidden" onClick={logoutOnClick}>
                Log out
              </UnstyledButton>
            </>
          )}
        </Group>
      </Flex>
      <div className="min-h-[calc(100vh-4.5rem)]">{children}</div>
    </div>
  );
}

export default MainLayout;
