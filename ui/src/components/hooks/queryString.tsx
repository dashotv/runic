import { useCallback } from 'react';

export function useQueryString() {
  const queryString = useCallback(form => {
    return objectToQueryString(form);
  }, []);

  return { queryString };
}

export function objectToQueryString(obj) {
  return new URLSearchParams(obj).toString();
}
