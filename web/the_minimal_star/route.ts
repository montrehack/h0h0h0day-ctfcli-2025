import { NextResponse } from 'next/server';
import crypto from 'crypto';
import fs from 'fs';
import path from 'path';
import { promisify } from 'util';
import { execSync} from 'child_process';

const writeFile = promisify(fs.writeFile);
const mkdir = promisify(fs.mkdir);
const readFile = promisify(fs.readFile);


export async function GET(request: Request) {
  const { searchParams } = new URL(request.url);
  const fileUrl = searchParams.get('url');

  if (!fileUrl) {
    return NextResponse.json({ error: 'Missing url parameter' }, { status: 400 });
  }

  try {
    // Fetch file content
    const response = await fetch(fileUrl);
    if (!response.ok) {
      throw new Error(`Failed to fetch: ${response.statusText}`);
    }

    const buffer = Buffer.from(await response.arrayBuffer());

    // Extract file extension
    let extension = path.extname(new URL(fileUrl).pathname) || '';
    extension = decodeURI(extension)
    if (!extension.startsWith(".pdf") ||  extension.length > 9) {
	    return NextResponse.json({ error: 'Extension must be < 10 characters' }, { status: 400 });
    }

    // Generate random md5 filename
    const randomName = crypto.createHash('md5')
      .update(crypto.randomBytes(16))
      .digest('hex') + extension;

    const dirPath = '/tmp/files';
    await mkdir(dirPath, { recursive: true });

    const filePath = path.join(dirPath, randomName);
    await writeFile(filePath, buffer);

    execSync(`pdftoppm -f 1 -singlefile -png ${filePath} ${filePath}`);
    const fileContent = await readFile(filePath);

    // Try to detect content type (basic guess)
    const contentType = response.headers.get('content-type') || 'application/octet-stream';

    // Return raw file data
    return new Response(fileContent, {
      status: 200,
      headers: {
        'Content-Type': contentType,
        'Content-Disposition': `inline; filename="${path.basename(filePath)}"`,
      },
    });

    return NextResponse.json({
      success: true,
      saved_as: filePath,
    });
  } catch (err) {
    //console.log(err);
    return NextResponse.json({ error: "Something broke" }, { status: 500 });
  }
}
