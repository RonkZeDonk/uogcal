import { CSSProperties } from "react";
import { LinkProps, Link as RRLink } from "react-router-dom";

type Props = {
  children?: string | string[] | JSX.Element | JSX.Element[]
  style?: CSSProperties
} & LinkProps

function Link({ children, style, ...props }: Props) {
  if (!style) {
    style = {};
  }
  if (!style.color) {
    style.color = "var(--text-color)";
  }

  return (
    <RRLink style={style} {...props}>{children}</RRLink>
  );
}

export default Link;
