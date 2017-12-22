1. When run without parameters it should look for `.lwrc.yaml` in working directory.
1.1 If `.lwrc.yaml` not found, it should terminate with error: No settings found.
1.2 Else it should parse `.lwrc.yaml`
1.3 If parsing failed, it should terminate with parsing error.
1.4 Else it should validate required options and it's values
1.5 If options validation failed, it should terminate with validation errors.

2. When run with -V, --version, it should return it's version
3. When run with -h, --help, it should return usage information
4. When run with options it should validate required options and it's values

5. Supported cloud providers: AWS
6. Supported engines: node
7. Supported library sources: npm, git

Forkflow:
1. Read command line arguments or options file
2. Create `.lwtmp`, `.lwtmp/lib`, `.lwtmp/build` directories
3. Install/copy/clone library into `.lwtmp/lib` and `.lwtmp/build`
3.1 Install dependencies to `.lwtmp/lib` and production dependencies to `.lwtmp/build`
3.2 Run unit tests at `.lwtmp/lib`
4. Take `index.js` template and inject library
4.1 Inject services to template
5. Zip package for deploy

!!! Supported services engines and lib sources should be validated in concrete implementation.
