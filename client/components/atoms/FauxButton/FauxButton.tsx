interface Props {
  label: string;
  selected: boolean;
}

export const FauxButton = (props: Props) => {
  const { label, selected, ...rest } = props;

  return (
    <div className="flex items-center w-fit cursor-pointer">
      <div
        {...rest}
        className={`${
          selected
            ? 'bg-violet-600 text-violet-100 hover:bg-violet-700'
            : 'bg-violet-100 text-slate-900 hover:bg-white'
        } px-4 py-2 font-medium tracking-wide rounded-md w-fit duration-150 transition-colors ease-in-out`}
      >
        {label}
      </div>
    </div>
  );
};
