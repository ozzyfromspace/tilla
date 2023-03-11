import Button from '@/components/atoms/Button/Button';
import RowDatePicker from '@/components/atoms/RowDatePicker/RowDatePicker';
import { BASE_URL } from '@/constants';
import axios from 'axios';
import {
  Dispatch,
  FormEvent,
  MutableRefObject,
  SetStateAction,
  useRef,
  useState,
} from 'react';

function toBasicISO(date: Date): string {
  const parts = date.toISOString().split('T');
  const timepart = `${parts[1].split('.')[0]}Z`;
  return `${parts[0]}T${timepart}`;
}

const handleDownloadRequest =
  (
    startDate: Date,
    endDate: Date,
    setDynamicLink: Dispatch<SetStateAction<string>>,
    linkRef: MutableRefObject<HTMLAnchorElement | null>
  ) =>
  (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    interface Payload {
      minLocalTime: string;
      maxLocalTime: string;
    }

    interface GoResponse {
      filename: string;
      msg: string;
    }

    const payload: Payload = {
      minLocalTime: toBasicISO(startDate),
      maxLocalTime: toBasicISO(endDate),
    };

    axios.post<GoResponse>(`${BASE_URL}/excel`, payload).then((resp) => {
      if (resp.status === 201) {
        // celebrate!
        console.log(resp.data);
        setDynamicLink(() => `${BASE_URL}/excel/${resp.data.filename}`);
        setTimeout(() => {
          linkRef.current?.click();
        }, 10);
      }
    });
  };

const handleDateSelection = (
  targetDateSetter: Dispatch<SetStateAction<Date>>,
  refDateSetter: Dispatch<SetStateAction<Date>>,
  refDate: Date,
  startingDate: boolean
) => {
  return (newDate: Date | null) => {
    if (!newDate) return;

    if (startingDate) {
      const newDateCopy = new Date(newDate);
      newDateCopy.setHours(0, 0, 0, 0);

      if (refDate.getTime() <= newDateCopy.getTime()) {
        targetDateSetter(() => newDateCopy);
        refDateSetter(() => {
          const endDate = new Date(newDateCopy);
          endDate.setHours(23, 59, 59, 0);
          return endDate;
        });

        return;
      }

      targetDateSetter(() => newDate);
      return;
    }

    const newEndDate = new Date(newDate);
    newEndDate.setHours(23, 59, 59, 0);

    if (refDate.getTime() >= newEndDate.getTime()) {
      targetDateSetter(() => newEndDate);
      refDateSetter(() => {
        const startDate = new Date(newEndDate);
        startDate.setHours(0, 0, 0, 0);
        return startDate;
      });

      return;
    }

    targetDateSetter(() => newDate);
  };
};

const AddExcelDownloadForm = () => {
  const [dynamicLink, setDynamicLink] = useState(() => '');
  const linkRef = useRef<HTMLAnchorElement | null>(null);

  const [startDate, setStartDate] = useState(() => {
    const d = new Date();
    d.setDate(1);
    d.setHours(0, 0, 0, 0);
    return d;
  });

  const [endDate, setEndDate] = useState(() => {
    const d = new Date();
    d.setMonth(d.getMonth() + 1);
    d.setDate(0);
    d.setHours(23, 59, 59, 0);
    return d;
  });

  return (
    <div>
      <form
        onSubmit={handleDownloadRequest(
          startDate,
          endDate,
          setDynamicLink,
          linkRef
        )}
        className="flex gap-2 flex-wrap"
      >
        <RowDatePicker
          label="Start Date"
          selectedDate={startDate}
          setSelectedDate={handleDateSelection(
            setStartDate,
            setEndDate,
            endDate,
            true
          )}
        />
        <RowDatePicker
          label="End Date"
          selectedDate={endDate}
          setSelectedDate={handleDateSelection(
            setEndDate,
            setStartDate,
            startDate,
            false
          )}
        />
        <Button label="Download Excel File" />
      </form>
      <a href={dynamicLink} ref={linkRef}></a>
    </div>
  );
};

export default AddExcelDownloadForm;
