interface Props {
  title: string;
  subtitle: string;
}

const Header = (props: Props) => {
  const { title, subtitle } = props;

  return (
    <header className="my-16 space-y-6">
      <h1 className="text-center text-white font-semibold text-6xl">{title}</h1>
      <h2 className="text-center text-white font-medium text-xl">{subtitle}</h2>
    </header>
  );
};

export default Header;
