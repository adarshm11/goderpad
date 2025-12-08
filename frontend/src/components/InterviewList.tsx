import { useInterviews, type Interview } from '../hooks/useInterviews.ts';
import InterviewTile from './InterviewTile.tsx';

function InterviewList() {
  const { interviews, loading, error } = useInterviews();

  return (
    <div>
      <h1>Interview List</h1>
      {loading && <p>Loading...</p>}
      {error && <p>Error: {error}</p>}
      {interviews.map((interview: Interview) => (
        <InterviewTile key={interview.title} interview={interview} />
      ))}
    </div>
  )
}

export default InterviewList;