import React from 'react';

// スタイルの定義
const style = {
  backgroundColor: '#c6e5d9',
  width: '400px',
  height: '30px',
  padding: '8px',
  margin: '8px',
  borderRadius: '8px',
};

// propsの型定義
interface InputTodoProps {
  todoText: string;
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  onClick: () => void;
  disabled: boolean;
}

// InputTodoコンポーネント
export const InputTodo: React.FC<InputTodoProps> = ({ todoText, onChange, onClick, disabled }) => {
  return (
    <div style={style}>
      <input disabled={disabled} placeholder="Todoを入力" value={todoText} onChange={onChange} />
      <button disabled={disabled} onClick={onClick}>
        追加
      </button>
    </div>
  );
};
