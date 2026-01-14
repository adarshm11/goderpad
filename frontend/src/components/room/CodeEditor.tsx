import Editor from '@monaco-editor/react';
import { useState, useRef, useContext, useEffect } from 'react';
import { DarkModeContext, UserContext } from '../../App';
import { SandpackProvider, SandpackPreview } from '@codesandbox/sandpack-react';

interface CodeEditorProps {
  code: string;
  setCode: (code: string) => void;
  sendWsMessage: (message: any) => void;
  users: Array<{
    userId: string;
    userName: string;
    cursorPosition: {
      lineNumber: number;
      column: number 
    } | null
  }>;
}

function CodeEditor({ code, setCode, sendWsMessage, users }: CodeEditorProps) {
  const { isDark } = useContext(DarkModeContext);
  const { userId } = useContext(UserContext);
  const editorRef = useRef<any>(null);
  const monacoRef = useRef<any>(null);
  const decorationIdsRef = useRef<string[]>([]);
  const [sandpackKey, setSandpackKey] = useState(0);
  const [hasError, setHasError] = useState(false);
  const [debouncedCode, setDebouncedCode] = useState(code);

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
      sendWsMessage({
        userId,
        type: 'cursor_update',
        payload: {
          lineNumber: e.position.lineNumber,
          column: e.position.column
        }
      });
    });
  }

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      setCode(value);
      console.clear();
      sendWsMessage({
        userId,
        type: 'code_update',
        payload: {
          code: value
        }
      });
    }
  }

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

  useEffect(() => {
    if (!monacoRef.current || !editorRef.current) return;

    const newDecorations = users
      .filter(user => user.cursorPosition !== null)
      .map(user => ({
        range: new monacoRef.current.Range(
          user.cursorPosition!.lineNumber,
          user.cursorPosition!.column,
          user.cursorPosition!.lineNumber,
          user.cursorPosition!.column
        ),
        options: {
          className: 'remote-cursor',
          hoverMessage: { value: user.userName },
          beforeContentClassName: `remote-cursor-label`,
        }
      }));

    decorationIdsRef.current = editorRef.current.deltaDecorations(
      decorationIdsRef.current,
      newDecorations
    );
  }, [users]);

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
