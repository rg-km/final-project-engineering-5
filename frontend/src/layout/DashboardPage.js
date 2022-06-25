import { useEffect, useState } from 'react';
import ClipLoader from 'react-spinners/ClipLoader';
import BeasiswaDetail from '../components/BeasiswaDetail';
import BeasiswaList from '../components/BeasiswaList';
import { getSiswa } from '../lib/siswa';
import useAuthStore from '../store/auth';

function DashboardPage() {
  const [beasiswaList, setBeasiswaList] = useState([]);
  const [activeBeasiswa, setActiveBeasiswa] = useState(null);
  const user = useAuthStore((state) => state.user);

  useEffect(() => {
    getSiswa(user.token).then((res) => {
      setBeasiswaList(res.data);
    });
  }, []);

  return (
    <div>
      {!beasiswaList ? (
        <div className="mt-16 text-center">
          <ClipLoader />
        </div>
      ) : (
        <div className="mx-auto grid max-w-screen-lg grid-cols-[1fr_2fr] items-start gap-4 py-16 px-4">
          <BeasiswaList
            beasiswaList={beasiswaList}
            activeBeasiswa={activeBeasiswa}
            setActiveBeasiswa={setActiveBeasiswa}
          />
          <BeasiswaDetail beasiswa={activeBeasiswa} />
          {/* {detailUser.dataBeasiswa && (
            
            <TableSiswaComponent detailSiswa={detailUser} />
            <BeasiswaDetail/> */}
        </div>
      )}
    </div>
  );
}

export default DashboardPage;
