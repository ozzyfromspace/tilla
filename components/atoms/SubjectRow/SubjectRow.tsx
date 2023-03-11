import { ChangeEvent, useId } from 'react';

export interface SubjectRowProps {
  fieldName: string;
  price: string;
  id: string;
  rowIndex: number;
  handleDelete: (rowId: string) => void;
  handleFieldChange?: (
    rowId: string,
    fieldName: 'fieldName' | 'price'
  ) => (e: ChangeEvent<HTMLInputElement>) => void;
}

const SubjectRow = (props: SubjectRowProps) => {
  const { fieldName, price, handleDelete, rowIndex, handleFieldChange, id } =
    props;
  const _id = useId();
  const inputname_id = `inputname-${_id}`;
  const inputprice_id = `inputprice-${_id}`;

  return (
    <div className="flex justify-start items-center gap-3 p-3 bg-slate-100 text-slate-700 rounded-md max-w-lg">
      <span className="flex gap-1 w-4">
        <p>{rowIndex + 1}</p>
        <p>.</p>
      </span>
      <label htmlFor={inputname_id} className="sr-only">
        Enter subject
      </label>
      <input
        type="text"
        value={fieldName}
        onChange={(e) => {
          e.preventDefault();
          e.stopPropagation();

          handleFieldChange?.(id, 'fieldName')(e);
        }}
        id={inputname_id}
        placeholder="Subject"
        className="text-slate-900 w-2/3"
      />{' '}
      <div className="flex gap-1 w-1/3">
        <label htmlFor={inputprice_id} className="sr-only">
          Enter price
        </label>
        <span>
          <p>$</p>
        </span>
        <input
          type="text"
          value={price}
          onChange={handleFieldChange?.(id, 'price')}
          className="text-slate-900 w-full"
          id={inputprice_id}
          placeholder="Price"
        />
      </div>
      <button onClick={() => handleDelete(id)} aria-label="Delete Row">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={1.5}
          stroke="rgb(51, 65, 85)"
          className="w-6 h-6"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>
    </div>
  );
};

export default SubjectRow;
