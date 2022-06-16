import express from 'express';

const USERS = [
  {
    email: 'johndoe@email.com',
    role: 'SISWA',
    token:
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c',
  },
  {
    email: 'janedoe@email.com',
    role: 'MITRA',
    token:
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkphbmUgRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cMErWtEf7DxCXJl8C9q0L7ttkm-Ex54UWHsOCMGbtUc',
  },
];

const router = express.Router();

router.post('/', (req, res) => {
  const { email, password } = req.body;
  if (!password) {
    res.status(400).json({ message: 'Password is required' });
  }

  const user = USERS.find((user) => user.email === email);
  if (!user) {
    res.status(404).json({ message: 'User not found' });
  }
  res.json({ role: user.role, token: user.token });
});

export default router;
