import express from 'express';

const DUMMY_TOKEN =
  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkphbmUgRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cMErWtEf7DxCXJl8C9q0L7ttkm-Ex54UWHsOCMGbtUc';

const router = express.Router();

router.post('/signup', (_, res) => {
  res.json({
    role: 'MITRA',
    token: DUMMY_TOKEN,
  });
});

export default router;
