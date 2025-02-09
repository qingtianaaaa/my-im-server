from PIL import Image
import os
import sys

def compress_image(input_path, max_size_mb=1):
    # 打开图片
    img = Image.open(input_path)
    
    # 获取原始文件大小（字节）
    original_size = os.path.getsize(input_path)
    
    # 计算目标大小（字节）
    target_size = 1024 * 1024 * max_size_mb
    
    # 如果原始大小已经小于目标大小，直接返回
    if original_size <= target_size:
        print(f"图片已经小于 {max_size_mb}MB，无需压缩")
        return
    
    # 计算初始质量
    quality = 95
    output_path = input_path.rsplit('.', 1)[0] + '_compressed.png'
    
    while quality > 5:
        # 保存压缩后的图片
        img.save(output_path, 'PNG', quality=quality, optimize=True)
        
        # 检查压缩后的大小
        compressed_size = os.path.getsize(output_path)
        
        if compressed_size <= target_size:
            print(f"压缩成功！质量：{quality}%")
            print(f"原始大小: {original_size/1024/1024:.2f}MB")
            print(f"压缩后大小: {compressed_size/1024/1024:.2f}MB")
            print(f"输出文件: {output_path}")
            return
        
        # 如果还是太大，降低质量继续尝试
        quality -= 5

    print("警告：即使使用最低质量也无法达到目标大小")

def main():
    if len(sys.argv) < 2:
        print("使用方法: python compress_image.py <图片路径>")
        return
    
    input_path = sys.argv[1]
    if not os.path.exists(input_path):
        print("错误：文件不存在")
        return
    
    if not input_path.lower().endswith('.png'):
        print("错误：只支持PNG格式图片")
        return
    
    compress_image(input_path)

if __name__ == "__main__":
    main()