import jwt from 'jsonwebtoken';
import { JWT_SECRET } from '../db/user.js';

export function authenticate(req, res, next) {
  const authHeader = req.headers.authorization;

  if (authHeader) {
    const token = authHeader.split(' ')[1];

    jwt.verify(token, JWT_SECRET, (err, user) => {
      if (err) {
        res.sendStatus(403);
        return;
      }

      req.user = user;
      next();
    });
  } else {
    res.sendStatus(401);
  }
}

export function authenticateMitra(req, res, next) {
  if (req.user.role !== 'MITRA') {
    res.status(403).json({
      message: `User with role ${req.user.role} is not allowed to access this resource`,
    });
    return;
  }
  next();
}
