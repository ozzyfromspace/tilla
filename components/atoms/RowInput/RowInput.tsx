import { ChangeEvent, useId } from 'react';

interface Props {
  label: string;
  handleChange: (e: ChangeEvent<HTMLInputElement>) => void;
  inputValue: string;
}

const RowInput = (props: Props) => {
  const { label, inputValue, handleChange } = props;
  const id = useId();
  const inputId = `rowInput${id}`;

  return (
    <div>
      <label htmlFor={inputId}>{label}</label>
      <input
        type="text"
        id={inputId}
        onChange={handleChange}
        value={inputValue}
      />
    </div>
  );
};

export default RowInput;
