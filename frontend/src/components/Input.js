function Input({ name, value, type, label, onChange }) {
  return (
    <div>
      <label htmlFor={name} className="block text-sm font-medium">
        {label}
      </label>
      <input
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        id={name}
        className="mt-1 block w-full rounded-sm border-gray-400 py-1 px-2"
      />
    </div>
  );
}

export default Input;
