import express from 'express';

const DUMMY_TOKEN =
  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c';

const router = express.Router();

router.post('/signup', (_, res) => {
  res.json({
    role: 'SISWA',
    token: DUMMY_TOKEN,
  });
});

export default router;
