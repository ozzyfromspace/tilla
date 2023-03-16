interface Props {
  label: string;
  selected: boolean;
}

export const FauxButton = (props: Props) => {
  const { label, selected, ...rest } = props;

  return (
    <div className="flex items-center w-fit">
      <div
        {...rest}
        className={`${
          selected
            ? 'bg-violet-600 text-violet-100'
            : 'bg-violet-100 text-slate-900'
        } px-4 py-2 font-medium tracking-wide rounded-md w-fit duration-150 transition-colors ease-in-out`}
      >
        {label}
      </div>
    </div>
  );
};
