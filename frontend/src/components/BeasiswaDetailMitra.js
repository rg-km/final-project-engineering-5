import { format } from 'date-fns';
import { id } from 'date-fns/locale';
import { useEffect, useState } from 'react';
import { getBeasiswa } from '../lib/beasiswa';
import TableSiswaComponent from './TableSiswaComponent';

function BeasiswaDetailMitra({ idBeasiswa, listSiswa }) {
  const [beasiswa, setBeasiswa] = useState(null);

  useEffect(() => {
    if (idBeasiswa) {
      getBeasiswa(idBeasiswa).then((res) => {
        setBeasiswa(res.beasiswa);
      });
    }
  }, [idBeasiswa]);

  return (
    <div className="rounded-lg border border-gray-300 p-4 shadow-md">
      {!beasiswa ? (
        <p className="text-center">
          Pilih beasiswa untuk melihat info lebih detail
        </p>
      ) : (
        <>
          <h2 className="text-xl font-semibold">
            {beasiswa.judulBeasiswa || beasiswa.namaBeasiswa}
          </h2>
          {beasiswa.tanggalPembukaan && beasiswa.tanggalPenutupan && (
            <div className="mt-2 grid w-fit grid-cols-[auto_auto] gap-x-4">
              <span className="font-semibold">Buka</span>
              <time dateTime={beasiswa.tanggalPembukaan}>
                {format(new Date(beasiswa.tanggalPembukaan), 'dd MMMM yyyy', {
                  locale: id,
                })}
              </time>
              <span className="font-semibold">Tutup</span>
              <time dateTime={beasiswa.tanggalPenutupan}>
                {format(new Date(beasiswa.tanggalPenutupan), 'dd MMMM yyyy', {
                  locale: id,
                })}
              </time>
            </div>
          )}
          <p className="mt-4">{beasiswa.deskripsi}</p>
          <TableSiswaComponent dataSiswa={listSiswa} />
        </>
      )}
    </div>
  );
}

export default BeasiswaDetailMitra;
