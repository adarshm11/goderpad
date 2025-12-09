import { useState, useEffect } from 'react';

export interface Interview {
  id: string;
  title: string;
  content: string;
  interviewer: string;
  candidate: string;
}

export function useInterviews() {
  const [interviews, setInterviews] = useState<Interview[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;

    const fetchInterviews = async () => {
      try {
        setLoading(true);
        setError(null);

        // TODO: Replace with actual API call after MongoDB is set up
        const mockInterviews = await new Promise<Interview[]>((resolve) => {
          setTimeout(() => {
            resolve([
              { id: '1', title: 'Mock w/ Shimmy', content: '// foo', interviewer: 'Jason', candidate: 'Shimmy' },
              { id: '2', title: 'Mock w/ Freaknd', content: '// bar', interviewer: 'Jason', candidate: 'Freaknd' },
              { id: '3', title: 'Mock2 w/ Shimmy', content: '// baz', interviewer: 'Jason', candidate: 'Shimmy' },
            ]);
          }, 500);
        });

        if (!cancelled) {
          setInterviews(mockInterviews);
          setLoading(false);
        }
      } catch (err) {
        if (!cancelled) {
          setError(err instanceof Error ? err.message : 'Unknown error');
          setLoading(false);
        }
      }
    };

    fetchInterviews();

    return () => {
      cancelled = true;
    };
  }, []);

  return { interviews, loading, error };
}

