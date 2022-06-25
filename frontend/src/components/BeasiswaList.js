function BeasiswaList({ beasiswaList, activeBeasiswa, setActiveBeasiswa }) {
  return (
    <div className="overflow-hidden rounded-lg border-[1.5px] border-gray-300 shadow-md">
      {!beasiswaList
        ? 'Tidak terdapat beasiswa yang tersedia'
        : beasiswaList.map((beasiswa) => (
            <div
              className={`cursor-pointer border-t-[1.5px] border-gray-300 px-4 py-3 first:border-t-0 hover:bg-gray-200 ${
                beasiswa === activeBeasiswa ? 'bg-gray-200' : ''
              }`}
              key={beasiswa.id}
              onClick={() => {
                setActiveBeasiswa(beasiswa);
              }}
            >
              <p className="font-semibold leading-tight">
                {beasiswa.judulBeasiswa}
              </p>
              <p>{beasiswa.idMitra}</p>
            </div>
          ))}
    </div>
  );
}

export default BeasiswaList;
