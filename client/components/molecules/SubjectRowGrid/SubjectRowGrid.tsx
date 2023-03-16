import SubjectRow, {
  FieldName,
  SubjectRowProps,
} from '@/components/atoms/SubjectRow/SubjectRow';
import { ChangeEvent, Dispatch, SetStateAction, useCallback } from 'react';

interface Props {
  subjectRows: SubjectRowProps[];
  setSubjectRows: Dispatch<SetStateAction<SubjectRowProps[]>>;
}

const SubjectRowGrid = (props: Props) => {
  const { subjectRows, setSubjectRows } = props;

  const updateSubjectField = useCallback(
    function (rowId: string, fieldName: FieldName) {
      return (e: ChangeEvent<HTMLInputElement>) => {
        setSubjectRows((s) => {
          e.preventDefault();
          e.stopPropagation();

          const copy = JSON.parse(JSON.stringify(s)) as SubjectRowProps[];

          let foundIndex = -1;

          for (let index = 0; index < copy.length; index++) {
            const props = copy[index];

            if (props.id === rowId) {
              foundIndex = index;
              break;
            }
          }

          if (foundIndex == -1) return s;

          if (fieldName === 'course-name') {
            copy[foundIndex].courseName = e.target.value;
          }

          if (fieldName === 'price-per-session') {
            if (
              !copy[foundIndex].pricePerSession.includes('.') &&
              e.target.value.endsWith('.')
            ) {
              copy[foundIndex].pricePerSession = e.target.value;
              return copy;
            }

            const newPrice = parseFloat(e.target.value);

            if (!!e.target.value.length && isNaN(newPrice)) {
              return s;
            }

            copy[foundIndex].pricePerSession = isNaN(newPrice)
              ? ''
              : `${newPrice}`;
          }

          if (fieldName === 'session-length-in-minutes') {
            copy[foundIndex].sessionLength =
              `${parseInt(e.target.value)}` || '';
          }

          return copy;
        });
      };
    },
    [setSubjectRows]
  );

  const handleDelete = useCallback(
    (rowId: string) => {
      setSubjectRows((s) => {
        const copy = JSON.parse(JSON.stringify(s)) as SubjectRowProps[];

        let foundIndex = -1;

        for (let index = 0; index < copy.length; index++) {
          const props = copy[index];

          if (props.id === rowId) {
            foundIndex = index;
            break;
          }
        }

        if (foundIndex === -1) return s;

        copy.splice(foundIndex, 1);
        return copy;
      });
    },
    [setSubjectRows]
  );

  return (
    <div className="mt-3 space-y-3 max-h-56 overflow-auto w-full">
      {subjectRows.map(({ id, ...rest }, index) => {
        return (
          <SubjectRow
            key={id}
            {...rest}
            rowIndex={index}
            id={id}
            handleFieldChange={updateSubjectField}
            handleDelete={handleDelete}
          />
        );
      })}
    </div>
  );
};

export default SubjectRowGrid;
