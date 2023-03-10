import { useId } from 'react';
import DatePicker from 'react-datepicker';

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
    <div>
      <label htmlFor={datePickerId}>{label}</label>
      <DatePicker
        id={datePickerId}
        selected={selectedDate}
        onChange={(date) => setSelectedDate?.(date ?? new Date())}
      />
    </div>
  );
};

export default RowDatePicker;
