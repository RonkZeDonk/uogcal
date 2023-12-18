import AuthForm from "@/components/AuthForm";
import Link from "@/components/Link";
import MainLayout from "@/layouts/main";
import { Box, Center, Space, Title } from "@mantine/core";

function SignUpPage() {
  return (
    <MainLayout>
      <Center mih="inherit">
        <Box>
          <Title order={1}>Register</Title>
          <AuthForm type="register" />
          <Space h="2rem" />
          <Center>
            <Link to="/login">Already have an account?</Link>
          </Center>
        </Box>
      </Center>
    </MainLayout>
  );
}

export default SignUpPage;
