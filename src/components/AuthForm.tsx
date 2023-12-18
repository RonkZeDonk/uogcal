import { Box, Button, Center, Dialog, Flex, Input, Text } from "@mantine/core";
import { useState } from "react";
import { IconContext } from "react-icons";
import { FaDiscord, FaGoogle, FaGithub, FaSlack } from "react-icons/fa";
import Link from "./Link";
import { useGoogleLogin } from "@react-oauth/google";
import { useNavigate } from "react-router-dom";

type AuthType = "login" | "register";
type OAuthType = "google" | "disabled";

interface AccountFormProps {
  type: AuthType;
}

function AuthForm({ type }: AccountFormProps) {
  const [showDialog, setShowDialog] = useState(false);
  const [username, setUsername] = useState("");
  const shouldShowUserError =
    username && (username.length < 4 || username.length > 16);
  const [password, setPassword] = useState("");
  const shouldShowPassError =
    password && (password.length < 8 || password.length > 64);
  const actionText = type == "login" ? "Log In" : "Sign Up";
  const navigate = useNavigate();

  const authFunctions: Record<OAuthType, () => void> = {
    "google": useGoogleLogin({
      onSuccess: res => {
        fetch("/auth/google/" + type, {
          method: "POST",
          body: JSON.stringify({
            accessToken: res.access_token,
          }),
        }).then((r) => {
          if (!r.ok) {
            // TODO inform the user that they may already have an account
            setShowDialog(true);
            return;
          }

          if (type == "register") {
            navigate("/import");
          } else {
            navigate("/account");
          }
        });
      },
    }),
    "disabled": () => setShowDialog(true),
  };

  return (
    <>
      <Box w="16rem">
        <form
          action={"/auth/" + type}
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
              <Link to="" onClick={authFunctions["google"]}>
                <FaGoogle />
              </Link>
              <Link to="" onClick={authFunctions["disabled"]}>
                <FaGithub />
              </Link>
              <Link to="" onClick={authFunctions["disabled"]}>
                <FaDiscord />
              </Link>
              <Link to="" onClick={authFunctions["disabled"]}>
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
            There was a problem with this authentication method... Please try
            again or use another method
          </Dialog>
        </form>
      </Box>
    </>
  );
}

export default AuthForm;
