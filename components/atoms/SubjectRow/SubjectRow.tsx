import { ChangeEvent, useId } from 'react';

export type FieldName =
  | 'course-name'
  | 'price-per-session'
  | 'session-length-in-minutes';

export interface SubjectRowProps {
  courseName: string;
  pricePerSession: string;
  sessionLength: string;
  id: string;
  rowIndex: number;
  handleDelete: (rowId: string) => void;
  handleFieldChange?: (
    rowId: string,
    fieldName: FieldName
  ) => (e: ChangeEvent<HTMLInputElement>) => void;
}

const SubjectRow = (props: SubjectRowProps) => {
  const {
    courseName,
    pricePerSession,
    sessionLength,
    handleDelete,
    rowIndex,
    handleFieldChange,
    id,
  } = props;
  const _id = useId();
  const inputname_id = `inputname-${_id}`;
  const inputprice_id = `inputprice-${_id}`;

  return (
    <div className="flex justify-start items-center gap-3 p-3 bg-slate-100 text-slate-700 rounded-md">
      <span className="flex gap-1 w-4">
        <p>{rowIndex + 1}</p>
        <p>.</p>
      </span>
      <label htmlFor={inputname_id} className="sr-only">
        Enter subject
      </label>
      <input
        type="text"
        value={courseName}
        onChange={(e) => {
          e.preventDefault();
          e.stopPropagation();

          handleFieldChange?.(id, 'course-name')(e);
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
          <p>($/Session)</p>
        </span>
        <input
          type="text"
          value={pricePerSession}
          onChange={handleFieldChange?.(id, 'price-per-session')}
          className="text-slate-900 w-14"
          id={inputprice_id}
          placeholder="60"
        />
      </div>
      <div className="flex gap-1 w-1/3">
        <label htmlFor={inputprice_id} className="sr-only">
          Session Length in Minutes
        </label>
        <span>
          <p>(minutes/Session)</p>
        </span>
        <input
          type="number"
          min={0}
          value={sessionLength}
          onChange={handleFieldChange?.(id, 'session-length-in-minutes')}
          className="text-slate-900 w-14"
          id={inputprice_id}
          placeholder="Price per session"
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
