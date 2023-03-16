import { ChangeEvent, useId } from 'react';

interface Props {
  label: string;
  handleChange: (e: ChangeEvent<HTMLInputElement>) => void;
  inputValue: string;
  placeholder?: string;
}

const RowInput = (props: Props) => {
  const { label, inputValue, handleChange, placeholder } = props;
  const id = useId();
  const inputId = `rowInput${id}`;

  return (
    <div className="grid grid-cols-[6rem,3fr] gap-2 p-3 bg-slate-100 w-full max-w-4xl rounded-md">
      <label htmlFor={inputId} className="text-slate-700">
        {label}
      </label>
      <input
        type="text"
        id={inputId}
        onChange={handleChange}
        value={inputValue}
        className="text-slate-900 w-full bg-slate-50"
        placeholder={placeholder ?? 'Enter Value'}
      />
    </div>
  );
};

export default RowInput;
