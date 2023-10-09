import AuthForm from "@/components/AuthForm";
import Link from "@/components/Link";
import MainLayout from "@/layouts/main";
import { Box, Center, Space, Title } from "@mantine/core";

function LoginPage() {
  return (
    <MainLayout>
      <Center mih="inherit">
        <Box>
          <Title order={1}>Login</Title>
          <AuthForm actionText="Login" formAction="/auth/login" />
          <Space h="2rem" />
          <Center>
            <Link to="/signup">Don't have an account?</Link>
          </Center>
        </Box>
      </Center>
    </MainLayout>
  );
}

export default LoginPage;
