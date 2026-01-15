import Editor from '@monaco-editor/react';
import { useState, useRef, useContext, useEffect } from 'react';
import { DarkModeContext, UserContext } from '../../App';
import { SandpackProvider, SandpackPreview } from '@codesandbox/sandpack-react';

interface CodeEditorProps {
  code: string;
  setCode: (code: string) => void;
  ws: WebSocket | null;
  users: Array<{
    userId: string;
    userName: string;
    cursorPosition: {
      lineNumber: number;
      column: number 
    } | null
  }>;
}

function CodeEditor({ code, setCode, ws, users }: CodeEditorProps) {
  const { isDark } = useContext(DarkModeContext);
  const { userId } = useContext(UserContext);
  const editorRef = useRef<any>(null);
  const monacoRef = useRef<any>(null);
  const [sandpackKey, setSandpackKey] = useState(0);
  const [hasError, setHasError] = useState(false);
  const [debouncedCode, setDebouncedCode] = useState(code);
  const decorationsRef = useRef<string[]>([]);

  const handleEditorWillMount = (monaco: any) => {
    monaco.editor.defineTheme('slate-dark', {
      base: 'vs-dark',
      inherit: true,
      rules: [],
      colors: {
        'editor.background': '#0f172a',
      },
    });
  };

  const handleEditorDidMount = (editor: any, monaco: any) => {
    editorRef.current = editor;
    monacoRef.current = monaco;

    editor.onDidChangeCursorPosition((e: any) => {
      if (ws && ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({
          userId,
          type: 'cursor_update',
          payload: {
            lineNumber: e.position.lineNumber,
            column: e.position.column
          }
        }));
      }
    });
  }

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      setCode(value);
      console.clear();
      if (ws && ws.readyState === WebSocket.OPEN) {
        ws.send(JSON.stringify({
          userId,
          type: 'code_update',
          payload: {
            code: value
          }
        }));
      }
    }
  };

  const handleSandpackReload = () => {
    setHasError(false);
    setSandpackKey(prev => prev + 1);
  };
  
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedCode(code);
    }, 500);

    return () => clearTimeout(timer);
  }, [code]);

  // Add effect to update cursor decorations when users change
  useEffect(() => {
    if (!editorRef.current || !monacoRef.current) return;

    const editor = editorRef.current;
    const monaco = monacoRef.current;

    // Remove old decorations
    decorationsRef.current = editor.deltaDecorations(decorationsRef.current, []);

    // Create new decorations for other users' cursors
    const newDecorations = users
      .filter(user => user.userId !== userId && user.cursorPosition)
      .map(user => {
        const position = user.cursorPosition!;
        return {
          range: new monaco.Range(
            position.lineNumber,
            position.column,
            position.lineNumber,
            position.column
          ),
          options: {
            className: `cursor-${user.userId}`,
            beforeContentClassName: `cursor-label-${user.userId}`,
            stickiness: monaco.editor.TrackedRangeStickiness.NeverGrowsWhenTypingAtEdges,
            hoverMessage: { value: user.userName }
          }
        };
      });

    decorationsRef.current = editor.deltaDecorations([], newDecorations);

    // Add dynamic styles for each user's cursor
    users
      .filter(user => user.userId !== userId)
      .forEach((user, index) => {
        const colors = ['#FF6B6B', '#4ECDC4', '#45B7D1', '#FFA07A', '#98D8C8', '#F7DC6F'];
        const color = colors[index % colors.length];

        // Remove old style if exists
        const oldStyle = document.getElementById(`cursor-style-${user.userId}`);
        if (oldStyle) oldStyle.remove();

        // Add new style
        const style = document.createElement('style');
        style.id = `cursor-style-${user.userId}`;
        style.textContent = `
          .cursor-${user.userId} {
            border-left: 2px solid ${color} !important;
            animation: blink 1s infinite;
          }
          .cursor-label-${user.userId}::before {
            content: "${user.userName}";
            position: absolute;
            top: -18px;
            left: -2px;
            background: ${color};
            color: white;
            padding: 2px 6px;
            border-radius: 3px;
            font-size: 11px;
            font-weight: 500;
            white-space: nowrap;
            z-index: 10;
          }
          @keyframes blink {
            0%, 49% { opacity: 1; }
            50%, 100% { opacity: 0.3; }
          }
        `;
        document.head.appendChild(style);
      });

    // Cleanup function
    return () => {
      users.forEach(user => {
        const style = document.getElementById(`cursor-style-${user.userId}`);
        if (style) style.remove();
      });
    };
  }, [users, userId]);

  return (
    <div className={`flex flex-row gap-4 ${isDark ? 'bg-slate-900' : 'bg-gray-100'} p-6 pt-20`}>
      <div className={`border-2 ${isDark ? 'border-white' : 'border-gray-900'} rounded-lg overflow-hidden w-1/2`}>
        <Editor 
          height='85vh' 
          defaultLanguage={'javascript'} 
          value={code} 
          theme={isDark ? 'slate-dark' : 'vs'}
          beforeMount={handleEditorWillMount}
          onMount={handleEditorDidMount}
          onChange={handleEditorChange}
          options={{
            minimap: { enabled: false },
            fontSize: 14,
            padding: { top: 10 },
            scrollBeyondLastLine: false,
          }}
        />
      </div>
      <div className={`w-1/2 border-2 ${isDark ? 'border-white' : 'border-gray-900'} rounded-lg overflow-hidden h-[85vh] relative`}>
        {hasError && (
          <div className='absolute top-2 right-2 z-10'>
            <button 
              onClick={handleSandpackReload}
              className='bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded'
            >
              Reload Preview
            </button>
          </div>
        )}
        <SandpackProvider
          key={sandpackKey}
          template={'react'}
          files={{
            '/App.js': debouncedCode,
          }}
          theme={isDark ? 'dark' : 'light'}
          options={{
            externalResources: [],
            bundlerURL: 'https://sandpack-bundler.codesandbox.io',
            recompileMode: 'delayed',
            recompileDelay: 300,
          }}
          style={{ height: '100%' }}
        >
          <SandpackPreview 
            style={{ height: '100%' }}
            showOpenInCodeSandbox={false}
            showRefreshButton={true}
          />
        </SandpackProvider>
      </div>
    </div>
  );
}

export default CodeEditor;
