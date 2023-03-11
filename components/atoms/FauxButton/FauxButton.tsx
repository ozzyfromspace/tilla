export const FauxButton = (props: { label: string }) => {
  const { label, ...rest } = props;

  return (
    <div className="flex items-center w-fit">
      <div
        {...rest}
        className="bg-violet-600 px-4 py-2 text-violet-100 font-medium tracking-wide rounded-md w-fit"
      >
        {label}
      </div>
    </div>
  );
};
