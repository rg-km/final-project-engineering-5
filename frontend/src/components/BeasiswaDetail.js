function BeasiswaDetail({ beasiswa }) {
  return (
    <div>
      {!beasiswa ? (
        <p className="flex min-h-full items-center justify-center bg-gray-200">
          Pilih beasiswa untuk melihat info lebih detail
        </p>
      ) : (
        <div className="bg-gray-200 p-4">
          <h2 className="text-xl font-semibold">{beasiswa.nama}</h2>
          <p>{beasiswa.namaMitra}</p>
          <div>
            <div>
              <span className="font-semibold">Buka</span>:{' '}
              <time dateTime={beasiswa.tanggalPembukaan}>
                {beasiswa.tanggalPembukaan}
              </time>
            </div>
            <div>
              <span className="font-semibold">Tutup</span>:{' '}
              <time dateTime={beasiswa.tanggalPenutupan}>
                {beasiswa.tanggalPenutupan}
              </time>
            </div>
          </div>
          <p className="mt-4">{beasiswa.deskripsi}</p>
        </div>
      )}
    </div>
  );
}

export default BeasiswaDetail;
