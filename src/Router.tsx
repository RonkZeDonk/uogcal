// Regex and glob adapted from:
// https://omarelhawary.me/blog/file-based-routing-with-react-router/

import { RouteObject, createBrowserRouter } from "react-router-dom";
import Link from "./components/Link";

type PageType = {
  default?: () => JSX.Element;
} & RouteObject;
const pages = import.meta.glob<PageType>("/src/pages/**/[a-z[]*.tsx");

function buildRoutes() {
  const routes: RouteObject[] = Object.keys(pages).map((page) => {
    // Each replace does the following: (taken from the omarelhawary blog)
    //  - Remove /src/pages, index and .tsx from each path
    //  - Replace [...param] patterns with *
    //  - Replace [param] patterns with :param
    const path = page
      // changes from the blog:
      // - added a start of line directive (^) to this replace call
      //   - if there is a file in path /src/pages/src/pages/ it would have been ignored
      // - match / before and . after index match so that pages like "someindex" can exist
      .replace(/^\/src\/pages|(?:\/)index(?=\.)|\.tsx$/g, "")
      .replace(/\[\.{3}.+\]/, "*")
      .replace(/\[(.+)\]/, ":$1");

    return {
      path,
      lazy: async () => {
        const { default: Component, ...others } = await pages[page]();

        return {
          Component: Component || (() => <></>),
          ...others,
        };
      },
    };
  });

  routes.push({
    path: "*",
    element: (
      <>
        <h1>Page Not Found</h1>
        <Link to="/">{"<"} Go Home</Link>
      </>
    ),
  });

  return routes;
}

const Router = createBrowserRouter(buildRoutes());

export default Router;
