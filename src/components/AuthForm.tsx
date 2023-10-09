import { Box, Button, Center, Dialog, Flex, Input, Text } from "@mantine/core";
import { useState } from "react";
import { IconContext } from "react-icons";
import { FaDiscord, FaGoogle, FaGithub, FaSlack } from "react-icons/fa";
import Link from "./Link";
import { useGoogleLogin } from "@react-oauth/google";

interface AccountFormProps {
  actionText: string;
  formAction: string;
}

function AuthForm({ actionText, formAction }: AccountFormProps) {
  const [showDialog, setShowDialog] = useState(false);
  const [username, setUsername] = useState("");
  const shouldShowUserError =
    username && (username.length < 4 || username.length > 16);
  const [password, setPassword] = useState("");
  const shouldShowPassError =
    password && (password.length < 8 || password.length > 64);

  const googleLogin = useGoogleLogin({
    onSuccess: console.log,
  });

  return (
    <>
      <Box w="16rem">
        <form
          action={formAction}
          method="post"
          encType="multipart/form-data"
          className="[&>*]:my-2"
        >
          <Input.Wrapper label="Username" required w="100%">
            <Input
              name="Username"
              value={username || ""}
              error={shouldShowUserError}
              onChange={(e) => {
                setUsername(e.currentTarget.value);
              }}
            />
            {shouldShowUserError ? (
              <Input.Error>
                Username should be between 4 and 16 characters
              </Input.Error>
            ) : (
              <></>
            )}
          </Input.Wrapper>
          <Input.Wrapper label="Password" required w="100%">
            <Input
              name="Password"
              value={password}
              error={shouldShowPassError}
              type="password"
              onChange={(e) => {
                setPassword(e.currentTarget.value);
              }}
            />
            {shouldShowPassError ? (
              <Input.Error>
                Password should be between 8 and 64 characters
              </Input.Error>
            ) : (
              <></>
            )}
          </Input.Wrapper>

          <Center pt="sm">
            <Button type="submit" variant="filled" w="100%">
              {actionText}
            </Button>
          </Center>

          <Center>
            <Text>Or {actionText.toLocaleLowerCase()} with:</Text>
          </Center>
          <IconContext.Provider value={{ size: "28px" }}>
            <Flex justify="space-around" mx="lg" pt="xs" pos="relative">
              <Link to="" onClick={() => googleLogin()}>
                <FaGoogle />
              </Link>
              <Link to="" onClick={() => setShowDialog(true)}>
                <FaGithub />
              </Link>
              <Link to="" onClick={() => setShowDialog(true)}>
                <FaDiscord />
              </Link>
              <Link to="" onClick={() => setShowDialog(true)}>
                <FaSlack />
              </Link>
            </Flex>
          </IconContext.Provider>
          <Dialog
            opened={showDialog}
            shadow="md"
            withBorder
            withCloseButton
            onClose={() => setShowDialog(false)}
          >
            This authentication method is currently unavailable... Please try
            another
          </Dialog>
        </form>
      </Box>
    </>
  );
}

export default AuthForm;
