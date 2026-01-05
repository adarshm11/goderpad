import Editor from '@monaco-editor/react';
import { useState, useRef, useContext, useEffect } from 'react';
import { DarkModeContext } from '../../App';
import { SandpackProvider, SandpackPreview } from '@codesandbox/sandpack-react';

function CodeEditor() {
  const { isDark } = useContext(DarkModeContext);
  const editorRef = useRef<any>(null);
  const [code, setCode] = useState('function App() {\n  return (\n    <div>\n      <h1>Hello, World!</h1>\n    </div>\n  );\n}\nexport default App;');
  const [sandpackKey, setSandpackKey] = useState(0);
  const [hasError, setHasError] = useState(false);

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

  const handleEditorDidMount = (editor: any, _: any) => {
    editorRef.current = editor;
  }

  const [debouncedCode, setDebouncedCode] = useState(code);
  
  useEffect(() => {
    const timer = setTimeout(() => {
      setDebouncedCode(code);
    }, 500);

    return () => clearTimeout(timer);
  }, [code]);

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      setCode(value);
    }
  }

  const handleSandpackReload = () => {
    setHasError(false);
    setSandpackKey(prev => prev + 1);
  };

  return (
    <div className={`flex flex-row gap-4 ${isDark ? 'bg-slate-900' : 'bg-gray-100'} p-6 pt-20`}>
      <div className={`border-2 ${isDark ? 'border-white' : 'border-gray-900'} rounded-lg overflow-hidden w-1/2`}>
        <Editor 
          height="85vh" 
          defaultLanguage={"javascript"} 
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
          <div className="absolute top-2 right-2 z-10">
            <button 
              onClick={handleSandpackReload}
              className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded"
            >
              Reload Preview
            </button>
          </div>
        )}
        <SandpackProvider
          key={sandpackKey}
          template={"react"}
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
