// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Plant.Outputs
{

    [OutputType]
    public sealed class Container
    {
        public readonly Pulumi.Plant.ContainerBrightness? Brightness;
        public readonly string? Color;
        public readonly string? Material;
        public readonly Pulumi.Plant.ContainerSize Size;

        [OutputConstructor]
        private Container(
            Pulumi.Plant.ContainerBrightness? brightness,

            string? color,

            string? material,

            Pulumi.Plant.ContainerSize size)
        {
            Brightness = brightness;
            Color = color;
            Material = material;
            Size = size;
        }
    }
}