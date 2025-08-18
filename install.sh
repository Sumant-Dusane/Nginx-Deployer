set -e 

echo "🚀 Installing nginx-deployer CLI..."

URL="https://github.com/Sumant-Dusane/Nginx-Deployer/releases/download/v1.0.0-beta/nginx-deployer_linux_arm_64"

curl -fsSL -o nginx-deployer "$URL" && \
chmod +x nginx-deployer && \
sudo mv nginx-deployer /usr/local/bin/

echo "✅ Installation complete!"
echo "📖 Usage: nginx-deployer --help"