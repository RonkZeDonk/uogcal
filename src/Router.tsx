// Code adapted from:
// https://omarelhawary.me/blog/file-based-routing-with-react-router/

import { FC, Fragment } from "react";
import { RouteObject, createBrowserRouter } from "react-router-dom";
import Link from "./components/Link";

type ImportType = Record<string, { default: FC }>;
const ROUTES: ImportType = import.meta.glob("/src/pages/**/[a-z[]*.tsx", { eager: true });

const routes = Object.keys(ROUTES).map((route) => {
  // Each replace does the following: (taken from the omarelhawary blog)
  //  - Remove /src/pages, index and .tsx from each path
  //  - Replace [...param] patterns with *
  //  - Replace [param] patterns with :param
  const path = route
    // change from blog: added a start of line directive (^) to this replace call
    // if there is a file in path /src/pages/src/pages/ it would have been ignored
    .replace(/^\/src\/pages|index|\.tsx$/g, "")
    .replace(/\[\.{3}.+\]/, "*")
    .replace(/\[(.+)\]/, ":$1");

  return { path, element: ROUTES[route].default };
});

const buildRoutes = () => {
  const arr: RouteObject[] = [];

  for (const { path, element: Element = Fragment } of routes) {
    arr.push({ path, element: <Element /> });
  }

  arr.push({ path: "*", element: (
    <>
      <h1>Page Not Found</h1>
      <Link to="/">{"<"} Go Home</Link>
    </>
  ) });

  return arr;
};

const Router = createBrowserRouter(buildRoutes());

export default Router;
