import { type Interview } from '../hooks/useInterviews.ts';

interface InterviewTileProps {
  interview: Interview;
}

function InterviewTile({ interview }: InterviewTileProps) {
  return (
    <div style={{
      backgroundColor: '#ffffff',
      borderRadius: '8px',
      padding: '16px',
      marginBottom: '12px',
      boxShadow: '0 1px 3px rgba(0, 0, 0, 0.1)',
      position: 'relative',
      display: 'flex',
      flexDirection: 'column',
      gap: '8px'
    }}>
      {/* Title, date, and collapse icon */}
      <div style={{
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'flex-start'
      }}>
        <div style={{ flex: 1 }}>
          <h2 style={{
            margin: 0,
            fontSize: '16px',
            fontWeight: 600,
            color: '#333333',
            marginBottom: '4px'
          }}>
            {interview.title}
          </h2>
        </div>
        {/* Collapse icon */}
        <div style={{
          color: '#888888',
          cursor: 'pointer',
          fontSize: '14px'
        }}>
          ─
        </div>
      </div>

      {/* Bottom section with candidate info */}
      <div style={{
        display: 'flex',
        alignItems: 'center',
        marginTop: '8px'
      }}>
        {/* Candidate info */}
        <div style={{
          display: 'flex',
          alignItems: 'center',
          gap: '8px'
        }}>
          <div style={{
            width: '32px',
            height: '32px',
            borderRadius: '50%',
            backgroundColor: '#9333ea',
            color: '#ffffff',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontSize: '14px',
            fontWeight: 600
          }}>
            {interview.candidate.charAt(0).toUpperCase()}
          </div>
          <span style={{
            color: '#2563eb',
            fontSize: '14px',
            fontWeight: 500,
            cursor: 'pointer'
          }}>
            {interview.candidate}
          </span>
        </div>
      </div>
    </div>
  );
}

export default InterviewTile;
