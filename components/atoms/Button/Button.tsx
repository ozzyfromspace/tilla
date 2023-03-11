interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  label: string;
}

const Button = (props: Props) => {
  const { label, ...rest } = props;

  return (
    <div className="flex items-center w-fit">
      <button
        {...rest}
        className="bg-violet-600 px-4 py-2 text-violet-100 font-medium tracking-wide rounded-md w-fit"
      >
        {label}
      </button>
    </div>
  );
};

export default Button;
