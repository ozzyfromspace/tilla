import { useId } from 'react';
import DatePicker from 'react-datepicker';

import 'react-datepicker/dist/react-datepicker.css';

interface Props {
  label: string;
  selectedDate: Date;
  setSelectedDate?: (value: Date | null) => void;
}

const RowDatePicker = (props: Props) => {
  const { label, selectedDate, setSelectedDate } = props;
  const id = useId();
  const datePickerId = `date-picker-${id}`;

  return (
    <div className="grid grid-cols-[2fr,3fr] gap-2 p-3 bg-slate-100 max-w-sm rounded-md">
      <label htmlFor={datePickerId} className="text-slate-700 flex">
        {label}
      </label>
      <DatePicker
        id={datePickerId}
        selected={selectedDate}
        onChange={(date) => setSelectedDate?.(date ?? new Date())}
        className="text-slate-700 cursor-pointer"
      />
    </div>
  );
};

export default RowDatePicker;
