export function checkPassword(password, confirm) {
  if (password !== confirm) {
    alert('Password berbeda');
    return false;
  }
  return true;
}
