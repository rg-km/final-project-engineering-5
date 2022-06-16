import express from 'express';
import jwt from 'jsonwebtoken';
import { JWT_SECRET, users } from '../db/user.js';

const router = express.Router();

router.post('/', (req, res) => {
  const { email, password } = req.body;
  console.log(req.body);
  if (!email || !password) {
    res.status(400).json({ message: 'Both email and password is required' });
    return;
  }

  const user = users.find(
    (user) => user.email === email && user.password === password
  );

  if (!user) {
    res.status(400).json({ message: 'Username or password incorrect' });
  }
  const token = jwt.sign({ email, role: user.role }, JWT_SECRET);
  res.json({ role: user.role, token: token });
});

export default router;
