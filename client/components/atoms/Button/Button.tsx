interface Props extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  label: string;
  selected: boolean;
}

const Button = (props: Props) => {
  const { label, selected, ...rest } = props;

  return (
    <div className="flex items-center w-fit">
      <button
        {...rest}
        className={`${
          selected
            ? 'bg-green-600 text-green-50'
            : 'bg-green-50 text-slate-900 border-[1px] border-slate-300'
        } px-4 py-2 font-medium tracking-wide rounded-md w-fit`}
      >
        {label}
      </button>
    </div>
  );
};

export default Button;
