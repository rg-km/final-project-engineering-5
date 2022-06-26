import { useEffect, useMemo, useState } from 'react';
import ClipLoader from 'react-spinners/ClipLoader';
import BeasiswaDetailSiswa from '../components/BeasiswaDetailSiswa';
import BeasiswaDetailMitra from '../components/BeasiswaDetailMitra';
import BeasiswaList from '../components/BeasiswaList';
import { getListBeasiswaSiswa } from '../lib/mitra';
import { getSiswa } from '../lib/siswa';
import useAuthStore from '../store/auth';

function DashboardPage() {
  const [beasiswaList, setBeasiswaList] = useState(null);
  const [activeBeasiswa, setActiveBeasiswa] = useState(null);
  const user = useAuthStore((state) => state.user);

  const listSiswaByBeasiswa = useMemo(() => {
    if (!beasiswaList) return;
    const obj = {};
    for (const beasiswaSiswa of beasiswaList) {
      const { idBeasiswa, idSiswa, namaSiswa } = beasiswaSiswa;
      const siswa = { id: idSiswa, nama: namaSiswa };
      if (obj.hasOwnProperty(beasiswaSiswa.idBeasiswa)) {
        obj[idBeasiswa].push(siswa);
      } else {
        obj[idBeasiswa] = [siswa];
      }
    }
    return obj;
  }, [beasiswaList]);

  useEffect(() => {
    if (user.role === 'SISWA') {
      getSiswa(user.token).then((res) => {
        setBeasiswaList(res.data);
      });
    } else if (user.role === 'MITRA') {
      getListBeasiswaSiswa(user.token)
        .then((res) => {
          setBeasiswaList(res.data);
        })
        .catch((err) => {
          if (err.response.status === 404) {
            setBeasiswaList([]);
          }
        });
    }
  }, [user.role, user.token]);

  return (
    <div>
      {!beasiswaList ? (
        <div className="mt-16 text-center">
          <ClipLoader />
        </div>
      ) : beasiswaList.length <= 0 ? (
        <p className="text-center font-medium">
          {user.role === 'SISWA'
            ? 'Kamu belum mendaftar beasiswa'
            : user.role === 'MITRA'
            ? 'Tidak terdapat beasiswa'
            : null}
        </p>
      ) : (
        <div className="mx-auto grid max-w-screen-lg grid-cols-[1fr_2fr] items-start gap-4 py-16 px-4">
          <BeasiswaList
            beasiswaList={beasiswaList}
            activeBeasiswa={activeBeasiswa}
            setActiveBeasiswa={setActiveBeasiswa}
          />
          {user.role === 'SISWA' ? (
            <BeasiswaDetailSiswa beasiswa={activeBeasiswa} />
          ) : user.role === 'MITRA' ? (
            <BeasiswaDetailMitra
              idBeasiswa={activeBeasiswa?.idBeasiswa}
              listSiswa={listSiswaByBeasiswa[activeBeasiswa?.idBeasiswa]}
            />
          ) : null}
        </div>
      )}
    </div>
  );
}

export default DashboardPage;
