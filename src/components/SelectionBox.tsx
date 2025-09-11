import {
  useMantineTheme,
  Combobox,
  Loader,
  Select,
} from "@mantine/core";

type SelectionBoxProps = {
  label: string;
  onSelect?: (n: string | null) => void;
  placeholder: string;
  options: string[];
  disabled?: boolean;
  loading?: boolean;
  side?: "right" | "left"; // assuming centre by default
};

export const SelectionBox = ({
  label,
  placeholder,
  options,
  disabled,
  loading,
  side,
  onSelect,
}: SelectionBoxProps) => {
  const theme = useMantineTheme();
  const radius =
    typeof theme.defaultRadius === "number"
      ? theme.defaultRadius
      : theme.radius[theme.defaultRadius];
  const leftRadius = `${side === "left" ? radius : 0}`;
  const rightRadius = `${side === "right" ? radius : 0}`;
  const radiusStr = `${leftRadius} ${rightRadius} ${rightRadius} ${leftRadius}`;

  return (
    <Select
      label={label}
      placeholder={placeholder}
      data={options}
      searchable
      disabled={disabled || loading}
      radius={radiusStr}
      selectFirstOptionOnChange
      onChange={(v) => onSelect && onSelect(v)}
      rightSectionPointerEvents="none"
      rightSection={loading ? <Loader size={18} /> : <Combobox.Chevron />}
      checkIconPosition="right"
      nothingFoundMessage={`${label} not found :(`}
    />
  );
};
