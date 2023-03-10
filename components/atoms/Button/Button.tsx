interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  label: string;
}

const Button = (props: Props) => {
  const { label, ...rest } = props;

  return <button {...rest}>{label}</button>;
};

export default Button;
