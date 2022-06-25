import express from 'express';
import siswaRoutes from './siswa.js';
import mitraRoutes from './mitra.js';
import beasiswaRoutes from './beasiswa.js';
import loginRoutes from './login.js';

const router = express.Router();

router.use('/siswa', siswaRoutes);
router.use('/mitra', mitraRoutes);
router.use('/beasiswa', beasiswaRoutes);
router.use('/login', loginRoutes);

export default router;
