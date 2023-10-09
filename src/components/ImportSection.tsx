import { Title } from "@mantine/core";

interface ImportSectionProps {
  children: JSX.Element | JSX.Element[];
  title: string;
}

function ImportSection({ children, title }: ImportSectionProps) {
  return (
    <div className="px-6 py-4 my-4 rounded-md bg-[var(--mantine-color-gray-1)] dark:bg-[var(--mantine-color-dark-6)]">
      <Title order={2} size="1.2rem">
        {title}
      </Title>
      {children}
    </div>
  );
}

export default ImportSection;
