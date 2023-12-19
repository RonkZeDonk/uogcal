import { RouterProvider } from "react-router-dom";
import Router from "./Router";
import { MantineColorsTuple, MantineProvider, createTheme } from "@mantine/core";
import { GoogleOAuthProvider } from "@react-oauth/google";

const myColor: MantineColorsTuple = [
  "#f7ecff",
  "#e7d7fa",
  "#cbadee",
  "#ac81e3",
  "#935bd9",
  "#8343d4",
  "#7c36d2",
  "#6a29bb",
  "#5e24a8",
  "#511b94",
];

const mantineTheme = createTheme({
  primaryColor: "myColor",
  colors: {
    myColor,
  },
});

function App() {
  return (
    <>
      <MantineProvider theme={mantineTheme} defaultColorScheme="auto">
        <GoogleOAuthProvider clientId={import.meta.env.VITE_G_OAUTH_CLIENT}>
          <RouterProvider router={Router} />
        </GoogleOAuthProvider>
      </MantineProvider>
    </>
  );
}

export default App;
