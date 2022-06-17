function BeasiswaList({ beasiswaList, activeBeasiswa, setActiveBeasiswa }) {
  return (
    <div className="space-y-4 bg-gray-200 p-4">
      {beasiswaList.map((beasiswa) => (
        <div
          className={`cursor-pointer border border-gray-500 p-2 hover:bg-gray-300 ${
            beasiswa === activeBeasiswa ? 'bg-gray-300' : ''
          }`}
          key={beasiswa.nama}
          onClick={() => {
            setActiveBeasiswa(beasiswa);
          }}
        >
          <p className="font-semibold">{beasiswa.nama}</p>
          <p>{beasiswa.namaMitra}</p>
        </div>
      ))}
    </div>
  );
}

export default BeasiswaList;
