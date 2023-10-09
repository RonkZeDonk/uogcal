export function getAuthToken() {
  const cookies = document.cookie.split(/; ?/);
  for (const cookie of cookies) {
    const [key, value] = cookie.split("=");

    if (key === "uogcal_token") {
      return value;
    }
  }

  return false;
}

export function isLoggedIn() {
  return getAuthToken() !== false;
}

export function logout() {
  fetch("/auth/logout");
  document.cookie = "uogcal_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC;";
}
