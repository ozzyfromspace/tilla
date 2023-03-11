interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  label: string;
}

const Button = (props: Props) => {
  const { label, ...rest } = props;

  return (
    <div className="flex items-center w-fit">
      <button
        {...rest}
        className="bg-violet-600 px-4 py-2 text-slate-50 font-medium tracking-wide rounded-md min-w-fit flex-1"
      >
        {label}
      </button>
    </div>
  );
};

export default Button;
